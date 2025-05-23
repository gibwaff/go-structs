# ДЗ 
Требуется реализовать методы в `BruteForceRecursion.java`, и проверить их тестами `BruteForceRecursionTest.java`

Запуск тестов: 
*Кнопкой в IntelliJ*

![IntelliJ RUN button](https://i.imgur.com/uHwKybe.png)
* Или командой `mvn -Dtest=BruteForceRecursion test` в терминале
## 1. Порядок выпуска фич
★★☆☆☆

У продакт менеджера есть список фич, которые надо выкатить. 
ПМ никак не может определиться с порядком, в котором их выпускать в продакшн. 
Если бы у него перед глазами был список всех возможных вариантов, то дело бы пошло быстрее.

Надо облегчить задачу менеджеру и написать метод, который принимает на вход массив с номерами фич 
`[1, 2, 3]` и возвращает все возможные перестановки выхода фич в продакшн.

```java
List<List<Integer>> orderOfReleaseFeatures(int[] numbersOfFeatures)
```
## 2. Pizza day
★★★☆☆

Ура, пятница! Pizza day настал! Осталось только определить, кто будет выносить коробки из-под пиццы. 
“Счастливчик” определяется следующим образом: 
генерируется случайное число и кто последний определит является ли оно суммой степеней тройки
(например, `12 = 3^1 + 3^2` или `91 = 3^0 + 3^2 + 3^4` является, а `21` - нет), 
тот и убирается в этот раз. Последние полгода проигрывает Иван, но в этот он решил схитрить и 
обратился к команде программистов для написания программы, которая будет решать эту задачу.

Помоги Ивану реализовать метод, который принимает на вход целое число и определяет, является ли оно
суммой степеней тройки.

```java
boolean checkPowersOfThree(int number)
```

## 3. Секретарь - шпион
★★★★☆

Директор ООО “Современные технологии” готовится выпустить новый продукт, который по его мнению 
взорвет рынок технологий. Но в последнее время его секретарша Оксана стала вести себя подозрительно:
поменяла новый iPhone на кнопочный телефон и стала писать слишком много сообщений. Директор 
подозревает, что она хочет слить информацию конкурентам. С помощью камер наблюдения удалось получить
информацию о последовательности цифр, которые набирала Оксана. Но вся трудность в том, что в 
кнопочном телефоне цифра обозначает несколько букв.

Директор просит вас написать метод, который берёт на вход строку из последовательности цифр, 
например `“43”` и выводит все возможные сообщения, которые могла написать Оксана.

```java
List<String> possibleMessages(String digits)
```
*Если вы забыли, как выглядит кнопочный телефон, то эта картинка для вас:*

![IntelliJ RUN button](https://i.imgur.com/1YzgsYB.png)