import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Main {

    private static int rotateDial(int dial, char direction, int value) {
        dial = (direction == 'L') ? dial - value : dial + value;
        dial = ((dial % 100) + 100) % 100; 
        return dial;
    }

    public static void main(String[] args) {
        if (args.length == 0) {
            System.err.println("Usage: java Main <input-file>");
            return;
        }

        int dial = 50;
        int count = 0;

        try (BufferedReader reader = new BufferedReader(new FileReader(args[0]))) {
            String line;
            while ((line = reader.readLine()) != null) {
                char dir = line.charAt(0);
                int val = Integer.parseInt(line.substring(1));
                dial = rotateDial(dial, dir, val);

                if (dial == 0) count++;
            }
        } catch (IOException e) {
            System.err.println("Error reading file: " + e.getMessage());
        }

        System.out.printf("Result: %d%n", count);
    }
}
