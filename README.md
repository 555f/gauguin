# Golang генератор HTTP транспорта REST и JSONRPC и многое другое

[![Go Reference][go-reference-badge]][go-reference]

> _**Примечание**_: На данный момент статус проекта в **BETA**. Пожалуйста, попробуйте и оставьте отзыв!
>  
> _**Примечание**_: Если вы считаете, что обнаружили критическую проблему в GG, пожалуйста со

> сообщите об этом, связавшись со мной по адресу

> [vitalylobchuk@gmail.com](mailto:vitalylobchuk@gmail.com).
>  

## Установка

```sh
git  clone  https://github.com/555f/gg.git
cd  gg
go  mod  download
make  install
```

или

```sh
go  install  github.com/555f/gg/cmd/gg
```

Добавьте `PATH="$HOME/go/bin:$PATH"` в `~/.bashrc`

## Начало

Создайте файл конфикурации `gg.yaml` в корне проекта.
В файле необзодимо установить пути в которых генератор будет искать теги:

```yaml
packages:
  -  ./internal/...
  -  ./pkg/...
```

так же в файле можно настравивать параметры плагинов:

```yaml
plugins:
  http:
    openapi-output:  ./docs
```

> _**Note**_: Все пути в файле указываються относительно файла конфигруации.
>  

Запуск генерации:

```shell
gg  run
```

## Теги

Теги - это параметры, которые участвуют в процессе генерации.
Cинтаксис: `// @<имя тега>:"<значение>,<опция1>,<опция1>,<параметр1=значение1>"`

Генерация работает на основе интерфейса либо структуры, для того чтобы включить генерацию необходимо добавить тег @gg:"<имя генератора>"

На данный момент в GG есть такие генераторы:

* http (генерация транспорта (клиента) REST, JSONRPC; openapi документации; запросов в формате CURL или HTTP; )

* klog (генерация мидлвары логирования с исрользованием `https://github.com/go-kit/log`)

* slog (генерация мидлвары логирования с исрользованием `https://pkg.go.dev/golang.org/x/exp/slog`)

* middleware (генерация хелперов для мидлвары)

* config (генерация загрузки настроек из env)

## http

> _**Примечание**_: По умолчанию генератор оборачивает json и xml ответы в имена возвращаемых значений метода

**Теги интерфейса**:

| Тег | Допустимые значения | Параметры | Описание |
| ---- | ---- | ---- | ---- |
| http-server |  |  | включить генерирацию сервера |
| http-type | echo, chi, mux |  | библиотека генерируемого сервера |
| http-client |  |  | влючить генерацию клиента |
| http-openapi |  |  | включить генерацию openapi документации |
| http-openapi-tags |  |  | теги для openapi документации применяемые ко всем методам интерфейса |
| http-openapi-header | имя HTTP заголовка | title | HTTP заголовок для openapi документации (например если заголовок передаются в через мидлвару в контексте) |
| http-api-doc |  |  | включить генерацию документации HTML |
| http-error | путь импорта до структуры ошибки |  | ошибка которую возвращает метод, используется для openapi документации и генерации клиента |
| http-req | http, curl |  | генерировать примеры запросов |

**Теги метода:**

| Тег | Допустимые значения | Опции | Параметры | Описание |
| ---- | ---- | ---- | ---- | ---- |
| http-method | GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH |  |  | Методы запроса |
| http-path |  |  |  | HTTP путь, для использования значений в пути надо добавить `:name` где имя это имя параметра метода значение которого должно быть передано в пути |
| http-openapi-tags | имена тегов через заяптую |  |  | теги openapi для метода |
| http-openapi-header | имя HTTP заголовка | required | title | HTTP заголовок для openapi документации (например если заголовок передаются в через мидлвару в контексте) |
| http-content-types | json, xml, urlencoded, multipart |  |  | тип HTTP контента |
| http-query-value | имя и значение query параметра через запятую |  |  | значения для query (только для клиента) |
| http-wrap-response | путь разделенный через точку в который надо обернуть |  |  | оборачивать ответ (для типа контента json или xml) |
| http-nowrap-response |  |  |  | не оборачивать ответ в имя возвращаемого значения (только если значение в ответе одно) |
| http-nowrap-request |  |  |  | не оборачивать запрос в имя параметра (только если параметр один) |
| http-error |  |  |  | тоже самое, что и для интерфейса только для конекртеного метода |
| http-time-format |  |  |  | формат времени (все из пакета time) по умолчанию time.RFC3339  |


**Теги параметров методов:**

| Тег | Допустимые значения | Параметры | Описание |
| ---- | ---- | ---- | ---- |
| http-name |  |  | Имя параметра в HTTP запросе |
| http-type | path, cookie, query, header, body |  | Тип нахождения параметра |
| http-required |  |  | Помечает поле как обязательное |
| http-flat |  |  | Делает отправку структуры развернутую, а не в поле |

**Теги структур:**

| Тег | Допустимые значения | Параметры | Описание |
| ---- | ---- | ---- | ---- |
| http-error-interface |  |  | Аннотация метода интерфейса (пример: `Code() string`) используется для обработки ошибок в клиенте. |
## Простые примеры

*  [REST](examples/rest-service)
*  [JSONRPC](examples/jsonrpc-service)
