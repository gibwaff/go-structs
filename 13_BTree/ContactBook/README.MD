# ДЗ 
Требуется реализовать методы в `ContactBook.java`, и проверить их тестами `ContactBookTest.java`

Запуск тестов: 

*Кнопкой в IntelliJ*

![IntelliJ RUN button](https://i.imgur.com/uHwKybe.png)
* Или командой `mvn -Dtest=ContactBookTest test` в терминале

### 1. Адресная книга
★★☆☆☆

Вам нужно помочь реализовать класс для хранения адресной книги и поиска по ней.

Для этого потребуется реализовать несколько функций. 

```java
// Добавляет новый контакт в адресную книгу
void add(Contact s);
```

```java
// Возвращает true, если контакт с именем name есть в книге
boolean contains(String name);
```

```java
// Возвращает количество контактов, добавленных в адресную книгу, имена которых начинаются с pref
int countStartsWith(String pref);
```

```java
// Возвращает все контакты, имена которых начинаются с префикса pref
Contact[] startsWith(String pref);
```