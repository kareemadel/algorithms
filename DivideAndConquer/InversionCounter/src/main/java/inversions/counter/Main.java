package inversions.counter;

import java.io.*;
import java.net.URL;
import java.util.Enumeration;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws IOException {
        ClassLoader classLoader = Main.class.getClassLoader();
        File file = new File(classLoader.getResource("inversions/counter/IntegerArray.txt").getFile());
        Scanner scanner = new Scanner(file);
        MyArrayList<Integer> list = new MyArrayList<>();
        while (scanner.hasNextInt()) {
            list.add(scanner.nextInt());
        }
        list.sort((Integer op1, Integer op2) -> op1 > op2);
        System.out.println(list.getInversions());
    }
}
