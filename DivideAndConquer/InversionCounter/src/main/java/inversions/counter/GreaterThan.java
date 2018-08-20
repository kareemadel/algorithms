package inversions.counter;

@FunctionalInterface
public interface GreaterThan<E> {
    public boolean greaterThan(E op1, E op2);
}
