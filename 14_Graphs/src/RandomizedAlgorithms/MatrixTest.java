package RandomizedAlgorithms;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Arrays;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class MatrixTest {

    static Stream<Arguments> providerMatrixTest() {
        int[][] matrix = { {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}};
        int[][] targetMatrix = { {1, 1, 1, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}};

        int[][] matrix2 = { {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}};
        int[][] targetMatrix2 = { {0, 0, 1, 1}, {0, 0, 1, 1}, {0, 0, 1, 1}, {0, 0, 1, 1}};

        return Stream.of(
                Arguments.of(matrix, 0, 1, 1, targetMatrix),
                Arguments.of(matrix2, 0, 3, 1, targetMatrix2));
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

    @DisplayName("Print a matrix")
    @ParameterizedTest
    @MethodSource("providerMatrixTest")
    void printTest(int[][] image, int row, int col, int newColor, int[][] targetImage) {
        var actual = Matrix.paint(image, row, col, newColor);
        var expected = targetImage;

        assertTrue(Arrays.deepEquals(actual,expected));
    }

    @DisplayName("Find count of Islands")
    @ParameterizedTest
    @MethodSource("providerNumIslandsTest")
    void numIslandsTest(char[][] grid, int islands) {
        var actual = Matrix.numIslands(grid);
        var expected = islands;

        assertEquals(actual, expected);
    }
}
