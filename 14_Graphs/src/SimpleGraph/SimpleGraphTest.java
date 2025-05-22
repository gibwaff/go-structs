package SimpleGraph;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.*;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class SimpleGraphTest {

    static String BFS(Node start) {
        HashSet<Node> seen = new HashSet();
        Queue<Node> queue = new LinkedList<>();
        String result = "";

        seen.add(start);
        queue.add(start);

        while (!queue.isEmpty()) {
            var current = queue.poll();
            result += current.val;

            Iterator<Node> i = current.neighbors.listIterator();
            while (i.hasNext()) {
                var next = i.next();
                if (!seen.contains(next)) {
                    seen.add(next);
                    queue.add(next);
                }
            }
        }
        return result;
    }

    static Stream<Arguments> providerGetImportanceTest() {
        var emp1 = new Employee(1, 5, new ArrayList<>(Arrays.asList(2, 3)));
        var emp2 = new Employee(2, 3, new ArrayList<>());
        var emp3 = new Employee(3, 3, new ArrayList<>());
        var emp4 = new Employee(4, 15, new ArrayList<>(Arrays.asList(5)));
        var emp5 = new Employee(5, 10, new ArrayList<>(Arrays.asList(6)));
        var emp6 = new Employee(6, 5, new ArrayList<>());

        return Stream.of(
                Arguments.of(new ArrayList<Employee>(Arrays.asList(emp1, emp2, emp3)), 1, 11),
                Arguments.of(new ArrayList<Employee>(Arrays.asList(emp4, emp5, emp6)), 4, 30)
        );
    }

    static Stream<Arguments> providerGraphTest() {
        var ak1 = new Node("Ivan");
        var ak2 = new Node("Maria");
        var ak3 = new Node("Kate");
        var ak4 = new Node("Peter");
        var ak5 = new Node("Anna");
        var ak6 = new Node("Emma");

        ak1.neighbors = new ArrayList(Arrays.asList(ak2, ak3));
        ak2.neighbors = new ArrayList(Arrays.asList(ak1, ak4));
        ak3.neighbors = new ArrayList(Arrays.asList(ak1, ak4));
        ak4.neighbors = new ArrayList(Arrays.asList(ak2, ak3, ak5, ak6));
        ak5.neighbors = new ArrayList(Arrays.asList(ak4));
        ak6.neighbors = new ArrayList(Arrays.asList(ak4));

        return Stream.of(
                Arguments.of(ak1),
                Arguments.of(ak2),
                Arguments.of(ak3),
                Arguments.of(ak4),
                Arguments.of(ak5),
                Arguments.of(ak6)
        );
    }

    @DisplayName("Get importance of employee")
    @ParameterizedTest
    @MethodSource("providerGetImportanceTest")
    void getImportanceTest(List<Employee> employees, int id, int importance) {
        var actual = SimpleGraph.getImportance(employees, id);
        var expected = importance;

        assertEquals(expected, actual);
    }

    @DisplayName("Help Ivan copy VK Graph")
    @ParameterizedTest
    @MethodSource("providerGraphTest")
    void cloneGraphVKTest(Node node) {
        var actual = BFS(SimpleGraph.cloneGraphVK(node));
        var expected = BFS(node);

        assertEquals(expected, actual);
    }
}
