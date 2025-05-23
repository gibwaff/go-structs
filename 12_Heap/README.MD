# ДЗ

Требуется реализовать методы в `Heap.java`, `WretchedPriorityQueue.java`, и проверить их тестами `HeapTest.java`, `WretchedPriorityQueueTest.java`

Запуск тестов:

*Кнопкой в IntelliJ*

![IntelliJ RUN button](https://i.imgur.com/uHwKybe.png)

* Или командой `mvn -Dtest=HeapTest test` в терминале

### 1. Построение кучи

★★★☆☆

В бухгалтерии вашей транспортной компании была установлена новая система обработки платежей разработанная компанией
Kriv and Rukov Ltd. Особенностью системы является то, что она принимает чеки по определенному принципу: на вход система 
принимает массив. В котором чек, с самой большой суммой находится на позиции 0 этого массива, два следующие по величине 
находятся на позициях 1 и 2 и так далее. Ваша система выгружает суммы чеков из бд в произвольном порядке. Ваша задача
реализовать метод, который на основании полученного массива сумм построит массив по правилам, описанным выше.

```java
 public static void buildHeapFromArray(int[] checks)
```


### 2. Поиск ошибки

★★★★☆

После того как вы смогли реализовать предыдущий метод и в программу стали поступать данные в нужном формате, обнаружилось, 
что в новой программе используется своя реализация очереди с приоритетами. И она работает с ошибками - ваш коллега написал
тесты `WretchedPriorityQueueTest.java` и два из них не проходит. Надо найти в коде `WretchedPriorityQueue.java` ошибку
и исправить ее, чтобы выполнялся один из тестов. А так же переписать код, чтобы методы добавления и удаления стали выполняться
быстрее и проходил тест по времени.

### 3. Поиск ближайших грузовиков

★★★☆☆

В вашей логистической компании хотят, чтобы вы реализовали метод поиска ближайших грузовиков к клиенту. Суть задачи такова:
вам приходит массив координат грузовиков и количество грузовиков, которое необходимо найти. Вы должны вернуть координаты ближайших к
клиенту грузовиков (координаты клиента считать [0, 0]), для этого реализуйте метод:

```java
public static List<TruckCoordinate> kClosestTrucks(List<TruckCoordinate> truckCoordinateList, int k)
```

### 4. Обработка фур

★★★★★

Кроме поиска ближайших грузовиков ваш босс хочет добавить возможность отслеживать график обработки грузов.
У вашей компании есть несколько складов 1,2,3 и так далее, их n штук.
Ежедневно на разгрузку приезжает несколько фур с товарами. В каждой фуре разное количество товара, поэтому каждая фура разгружается
за разное время t.
На вход в ваш метод поступает количество складов и массив чисел — время на разгрузку каждой фуры.
Фуры подъезжают на разгрузку к первому освободившемуся складу. Если несколько складов готовы принять автомобиль на разгрузку, то 
водитель подъезжает к самому ближнему - с минимальным номером.

Вернуть метод должен список, в котором отмечены времена, когда началась разгрузка каждого груза.

Допустим у нас один склад. Первая фура начала разгружаться, то есть у нее время разгрузки ноль. Она разгружалась 5 у.е. времени,
значит следующая фура начнет разгружаться только в 5. Тогда ответ [0, 5]
Если бы было больше складов, то обе фуры могли бы начать разгружаться в 0. В данном случае ответ [0,0]
```java
public static List<Long> unloadingTruck(int n, int[] times)
```


