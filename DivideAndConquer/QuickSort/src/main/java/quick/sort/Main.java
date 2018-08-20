package quick.sort;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Comparator;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) throws FileNotFoundException {
        ClassLoader classLoader = Main.class.getClassLoader();
        File file = new File(classLoader.getResource("quick/sort/IntegerList.txt").getFile());
        Scanner scanner = new Scanner(file);
        QuickSortArrayList<Integer> list = new QuickSortArrayList<>();
        while (scanner.hasNextInt()) {
            list.add(scanner.nextInt());
        }
        System.out.println("The number of comparisons");
        System.out.println(list.qSort((Comparator<Integer>) (o1, o2) -> o1 - o2));
    }
}
