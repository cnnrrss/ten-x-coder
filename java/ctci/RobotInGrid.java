import java.awt.Point;
import java.util.ArrayList;

public class RobotInGrid {
    
    public boolean pathExists;

    public RobotInGrid() {}

    public ArrayList<Point> getPathRecursive(boolean[][] maze) {
        if (maze == null || maze.length == 0) return null;
        ArrayList<Point> path = new ArrayList<Point>();
        if (getPathRecursive(maze, maze.length - 1, maze[0].length - 1, path)) {
            return path;
        }
        return null;
    }
    public boolean getPathRecursive(boolean[][] maze, int row, int col, ArrayList<Point> path) {
        /* If out of bounds, return */
        if (col < 0 || row < 0 || !maze[row][col]) {
            return false;
        }

        boolean isAtOrigin = (row == 0) && (col == 0);

        if (isAtOrigin || getPathRecursive(maze, row, col - 1, path) ||
                getPathRecursive(maze, row - 1, col, path)) {
                Point p = new getPathRecursive(row, col);
                path.add(p);
                return true;
            }
        return false;
    }

    public static void main(String[] args) {
        // TODO: write tests
        // RobotInGrid robot = new RobotInGrid();
        // Boolean[][] maze 
        // robot.getPathRecursive(maze);
    }
}