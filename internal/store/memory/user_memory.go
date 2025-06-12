package memory

import (
	"ApiTrain/internal/domain"
	"ApiTrain/internal/store/postgres/userrepo"
	"sync"
)

type MemoryUserRepo struct {
	mu    sync.RWMutex         // почему именно RW а не просто sync.Mutex?
	users map[int]*domain.User //есть идея ключем сделать не айдишик а например имя пользователя для быстрой проверки уникальности так как все ключи мапы должны быть уникальны но я хз насколько это чисто
	// просто мне каежтся чем меньше таких не явных штук тем лучше в палне понимания того что происходит и просто это делает код менее зависимым и гибким но мб и норм идея так сделать УЗНАТЬ У ЮЛИ
	nextId int
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{
		users:  make(map[int]*domain.User),
		nextId: 1,
	}
}
func (m *MemoryUserRepo) Create(user *domain.User) (*domain.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	user.Id = m.nextId //тут добавить мб проверку но существования пользователя так говорит чат гпт но я ему не верю и хочу просто сделать это в сервисе то есть тупо отдельным запросом проверять а не на уровне бд прямо при создании
	// хоть это и более оптимальный путь но бляха чуйка подсказывает что в работе за такое по рукам дадут
	m.users[user.Id] = user
	m.nextId++
	return user, nil
}
func (m *MemoryUserRepo) GetByUsername(username string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, exists := range m.users {
		if exists.Username == username {
			return true, nil
		}
	}
	return false, nil // тут перепроверить логику мб еще можно прикрутить сортировку (алгоритм для оптимизации поиска)
}
func (m *MemoryUserRepo) LoginByUsername(username string) (*domain.LoginUserInternal, error) {
	m.mu.Lock()                            // мне сказали про RLock я хз что это такое нужно поискать теорию
	var userLogin domain.LoginUserInternal //ранее уже практиковал анонимные структуры но так и не выяснил уместны ли они поэтому пока так
	defer m.mu.Unlock()
	for _, exists := range m.users {
		if exists.Username == username {
			userLogin.Id = exists.Id
			userLogin.PasswordHash = exists.Password // тут должен быть захешированный пароль не помню точно ли он там перепроверить
			userLogin.Username = exists.Username
			return &userLogin, nil
		}
	}
	return nil, userrepo.ErrUserNotFound
}
