package Courses;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class CoursesTest {

    static Stream<Arguments> providerCanFinishTest() {
        int[][] prerequisites1 = { {1, 0}};
        int[][] prerequisites2 = { {1, 0}, {0, 1}};
        int[][] prerequisites3 = { {4, 3}, {3, 2}, {2, 1}, {1, 0}};

        return Stream.of(
                Arguments.of(2, prerequisites1, true),
                Arguments.of(2, prerequisites2, false),
                Arguments.of(5, prerequisites3, true));
    }

    static Stream<Arguments> providerNumIslandsTest(){
        char[][] grid1 = {
                {'1', '1', '1', '1', '0'},
                {'1', '1', '0', '1', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '0', '0', '0'}
        };

        char[][] grid2 = {
                {'1', '1', '0', '0', '0'},
                {'1', '1', '0', '0', '0'},
                {'0', '0', '1', '0', '0'},
                {'0', '0', '0', '1', '1'}
        };

        return Stream.of(
                Arguments.of(grid1, 1),
                Arguments.of(grid2, 3));
    }

    @DisplayName("Can you finish all Skillbox cources")
    @ParameterizedTest
    @MethodSource("providerCanFinishTest")
    void canFinishTest(int numCourses, int[][] prerequisites, boolean result) {
        var actual = Courses.canFinish(numCourses, prerequisites);
        var expected = result;

        assertEquals(actual,expected);
    }
}
