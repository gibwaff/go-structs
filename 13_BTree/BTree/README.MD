# ДЗ 
Требуется реализовать методы в `BTree.java`, и проверить их тестами `BTreeTest.java`

Запуск тестов: 

*Кнопкой в IntelliJ*

![IntelliJ RUN button](https://i.imgur.com/uHwKybe.png)
* Или командой `mvn -Dtest=BTreeTest test` в терминале

### 1. BTree
★★★★★

Вам нужно реализовать B-дерево с добавлением элементов. Удаление не требуется. 

На его основе нужно реализовать несколько функций. 

Не забудьте проверить, что глубина у дерева получается такой, как вы ожидаете и не растет линейно от количества элементов. 
```java
// Добавляет число X в дерево
public void add(int x);
```
```java
// Проверяет, было ли число X добавлено в дерево
public boolean contains(int x)
```
```java
// Выводит текущую максимальную глубину дерева
public int getMaxHeight();
```

```java
// Возвращает список всех чисел, добавленных в дерево в возрастающем порядке
public int[] getSorted();
```