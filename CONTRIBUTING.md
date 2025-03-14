## Архитектура

Проект использует Clean Architecture с следующими слоями:

- `cmd/` - точки входа приложения
  - `tg_bot/` - Telegram бот
  - `print_json/` - утилита для печати конфигурации
  - `pprof/` - профилирование приложения

- `internal/` - внутренняя логика
  - `adapters/` - адаптеры для внешних сервисов (Telegram)
  - `configs/` - конфигурация приложения
  - `services/` - бизнес-логика
  - `telebot/routes/` - обработчики Telegram сообщений
  - `use_case/` - сценарии использования

- `pkg/` - публичные пакеты
  - `json/` - работа с JSON конфигурацией

## Make команды

```bash
# Запуск
make run APP_TELEGRAM_TOKEN=your_token_here    # Запуск Telegram бота
make run_print_json                            # Запуск утилиты печати конфигурации

# Тестирование
make test              # Запуск тестов
make test_coverage     # Запуск тестов с отчетом о покрытии

# Разработка
make mockgen           # Генерация моков (требуется go install go.uber.org/mock/mockgen@latest)
make lint              # Запуск линтера (golangci-lint)

# Профилирование
make run_pprof         # Запуск профилирования
make pprof_mem         # Анализ памяти
make pprof_cpu         # Анализ CPU

# Сборка и деплой
make build             # Сборка для Linux
make docker_build      # Сборка Docker образа
make docker_run_img APP_TELEGRAM_TOKEN=your_token_here    # Запуск в Docker
make docker_push       # Публикация образа в ghcr.io
```

## Переменные окружения

Для работы приложения требуются следующие переменные окружения:

- `APP_TELEGRAM_TOKEN` - токен Telegram бота, полученный от @BotFather

Вы можете передать переменную непосредственно при выполнении команды:
```bash
make run APP_TELEGRAM_TOKEN=your_token_here
```

Или установить её перед запуском:
```bash
export APP_TELEGRAM_TOKEN=your_token_here
make run
```

## Зависимости

- [fx](https://github.com/uber-go/fx) - DI контейнер
- [telebot](https://github.com/tucnak/telebot) - Telegram Bot API клиент
- [mock](https://github.com/uber-go/mock) - генерация моков для тестирования

### Мокирование

Для генерации моков используется `go.uber.org/mock`. Чтобы сгенерировать моки:

1. Добавьте комментарий перед интерфейсом:
```go
//go:generate mockgen -destination=../../mocks/services/mock_service.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services Service
type Service interface {
    DoSomething() error
}
```

2. Запустите генерацию:
```bash
make mockgen
```

Моки будут сгенерированы в директории `mocks/`. Пример использования в тестах:

```go
func TestService(t *testing.T) {
    mockController := gomock.NewController(t)
    defer mockController.Finish()

    mockService := services_mocks.NewMockService(mockController)
    mockService.EXPECT().DoSomething().Return(nil)
}
```

## Конфигурация

Бот использует JSON конфигурацию для определения реакций на сообщения:

```json
{
  // Общие настройки для всех регулярных выражений (опционально)
  "common": {
    // Префикс добавляется в начало каждого паттерна
    "prefix": "(?i)(\\A|\\s)(",
    // Суффикс добавляется в конец каждого паттерна
    "suffix": ")(\\s|\\z)"
  },

  // Повседневные реакции (массив, порядок важен - первое совпадение имеет приоритет)
  "everyday": [
    {
      // Регулярное выражение без префикса и суффикса
      "pattern": "привет|здравствуй",
      // Массив возможных реакций (одна будет выбрана случайно)
      "reactions": ["👍", "❤️"],
      // Переопределение префикса (опционально, наследуется из common)
      "prefix": "(?i)(",
      // Переопределение суффикса (опционально, наследуется из common)
      "suffix": ")"
    }
  ],

  // Праздничные периоды (ключ = название праздника)
  "holidays": {
    "new_year": {
      // Дата начала праздника
      "start_day": 25,
      "start_month": 12,
      // Дата окончания праздника
      "end_day": 5,
      "end_month": 1,
      // Праздничные реакции (заменяют обычные в период праздника)
      "reactions": ["🎄", "🎅"],
      // Специальные паттерны для праздника (опционально)
      // Сначала проверяются праздничные паттерны, если совпадений не найдено - проверяются everyday паттерны
      "patterns": [
        {
          "pattern": "с новым годом|рождество",
          "reactions": ["🎄", "🎅", "⛄"],
          // prefix и suffix можно переопределить для каждого паттерна
          "prefix": "(?i)(",  // опционально
          "suffix": ")"       // опционально
        }
      ]
    }
  }
}
``` 