# go-openai-client

# Основная информация

### Средства разработки
1) [**golang 1.23.0**](https://go.dev/doc/devel/release)
2) [**go-openai lib**](https://github.com/sashabaranov/go-openai/)

### Архитектура
| Наименование                                                                                   | Назначение | Тип      |
|------------------------------------------------------------------------------------------------|-------------|----------|
| <a name="core_name"></a>OpenAI Module (`openai`)                                                          |  Модуль   | lib      |

#### Первый запуск
1) Необходимо установить все указанные выше продукты настроить и заполнить необходимые данные
2) Настроить .env сервисов руководствуясь информацией из config.example.yml
3) Провести билды сервисов и ядра
4) Запустить сервисы с помощью скрипта runs.sh или с помощью docker
5) все базы данные и первичные данные баз данных заполнятся/обновятся автоматически


#### Заполнить config.yml следующими параметрами


    - gpt
```json
{
  gpt_api_key: '*************' # OpenAI API key for GPT integration
  gpt_api_url: 'https://api.openai.com/v1'
  gpt_assist_id: '*************' # Assistant ID for a specific GPT instance or model
  gpt_model: 'gpt-4o-mini'
}
```
