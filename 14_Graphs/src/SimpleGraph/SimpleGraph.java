package SimpleGraph;

import java.util.*;

public class SimpleGraph {

    // Task #1
    static int getImportance(List<Employee> employees, int id) {
        return -1; // Please implement
    }

    // Task #2
    static Node cloneGraphVK(Node node) {
        return null; // Please implement
    }
}

// Определение сотрудника.
class Employee {
    public int id;
    public int importance;
    public List<Integer> subordinates;

    public Employee(int id, int importance, List<Integer> subordinates){
        this.id = id;
        this.importance = importance;
        this.subordinates = subordinates;
    }
};

// Определение Node.
class Node {
    public String val;
    public List<Node> neighbors;
    public Node() {
        val = "";
        neighbors = new ArrayList<>();
    }
    public Node(String _val) {
        val = _val;
        neighbors = new ArrayList<>();
    }
    public Node(String _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
