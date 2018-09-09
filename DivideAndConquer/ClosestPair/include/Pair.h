#ifndef CLOSESTPAIR_PAIR_H
#define CLOSESTPAIR_PAIR_H
#include <Point.h>

class Pair {
    Point a, b;
    double distance;
    static Pair bruteForceClosestPair(Point P[], int n);
    static Pair findCloserPair(Pair pLeft, Pair pRight);
    static Pair stripClosestPair(Point strip[], int n, Pair pair);
    static Pair recursiveClosestPair(Point Px[], Point Py[], int n);

public:
    Pair();
    Pair(Point a, Point b);
    Pair(Point a, Point b, double distance);
    Point getFirstPoint() const;
    Point getSecPoint() const;
    double getDistance() const;
    static Pair findClosestPair(Point P[], int n);
};

#endif //CLOSESTPAIR_PAIR_H
