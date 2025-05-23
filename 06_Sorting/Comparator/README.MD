# ДЗ 
Требуется реализовать методы в `BookComparator.java`, и проверить их тестами `BookCompparatorTest.java`

Запуск тестов: 

*Кнопкой в IntelliJ*

![IntelliJ RUN button](https://i.imgur.com/uHwKybe.png)
* Или командой `mvn -Dtest=BookComparatorTest test` в терминале

### 1. Книжный магазин
★★☆☆☆

В классе `Book.java` описаны основные свойства книги - автор, название, год издания и цена. 
Нам нужно написать компаратор, который помог бы отсортировать их по цене.
В классе `BookPriceComparator.java` реализуйте метод
```java
public static int compareBooksByPrice(Book first, Book second)
```

Верните `int` равный 1 если цена списания текущей книги больше.

Верните `int` равный 0 если цены равны.

Верните `int` равный -1 если цена списания текущей книги меньше.


### 2. Сортировка на полке
★★★☆☆

У нас есть немного книг и мы хотели бы их отсортировать на своей полке по автору в лексикографическом порядке.
Если автор совпал, то по названию книги. Если же и название одинаковое, то по году издания.

```java
public static int compareBooksByAuthorAndTitleAndYear(Book first, Book second)
```

Верните `int` -1 если первая книга должна стоять раньше второй, иначе 1.

*Совет:*
Для сравнения строк пользуйтесь методом `compareTo` класса String.

*Например*
```java
"Грокаем алгоритмы", "Адитья Бхаргава", 2017
//идет раньше чем книга
"Грокаем алгоритмы", "Адитья Бхаргава", 2021
```        


### 3. Ищем ошибку в классе Book
★★★☆☆

В классе `Book.java` нам кто-то уже любезно реализовал метод `public int compareTo(Book o)`.
Но он работает неправильно — когда пользователи пытаются добавить в свою коллекцию на основе TreeSet 
некоторые книги туда не попадают. Исправьте метод так, чтобы сравнивались все поля..
Исправьте метод `public int compareTo(Book o)`, чтобы пользователь мог добавить все книги в свою коллекцию.