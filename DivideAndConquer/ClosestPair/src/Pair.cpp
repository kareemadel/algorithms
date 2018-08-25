#include <Point.h>
#include <Pair.h>
#include <stdlib.h>


Pair Pair::bruteForceClosestPair(Point P[], int n) {
    Pair closestPair = Pair(P[0], P[1]);
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            double dist = Point::distance(P[i], P[j]);
            if (dist < closestPair.distance) {
                closestPair = Pair(P[i], P[j], dist);
            }
        }
    }
    return closestPair;
}

Pair Pair::findCloserPair(Pair pLeft, Pair pRight) {
    if (pLeft.distance < pRight.distance) {
        return pLeft;
    } else {
        return pRight;
    }
}

Pair Pair::stripClosestPair(Point strip[], int n, Pair pair) {
    Pair closestPair = pair;
    for (int i = 0; i < n; ++i) {
        for (int j = i + 1; j < n && j <= i + 7; ++j) {
            double dist = Point::distance(strip[i], strip[j]);
            if (dist < closestPair.distance) {
                closestPair = Pair(strip[i], strip[j], dist);
            }
        }
    }
    return closestPair;
}

Pair Pair::recursiveClosestPair(Point Px[], Point Py[], int n) {
    if (n <= 3) {
        return bruteForceClosestPair(Py, n);
    }
    int halfSize = n / 2;
    Point midPoint = Px[halfSize - 1];
    int leftSize = halfSize, rightSize = n - halfSize;
    Point Pyl[leftSize], Pyr[rightSize];
    for (int i = 0, li = 0, lr = 0; i < n; ++i) {
        if (midPoint.compareX(Py[i]) >= 0 && li < leftSize) {
            Pyl[li++] = Py[i];
        } else {
            Pyr[lr++] = Py[i];
        }
    }
    Pair leftPair = recursiveClosestPair(Px, Pyl, halfSize);
    Pair rightPair = recursiveClosestPair(Px + halfSize, Pyr, n - halfSize);
    Pair closerPair = findCloserPair(leftPair, rightPair);

    Point strip[n];
    int stripSize = 0;
    for (int i = 0; i < n; ++i) {
        if (abs(midPoint.getX() - Py[i].getX()) < closerPair.distance) {
            strip[stripSize++] = Py[i];
        }
    }
    return stripClosestPair(strip, stripSize, closerPair);
}

Pair Pair::findClosestPair(Point P[], int n) {
    Point Px[n];
    Point Py[n];
    for (int i = 0; i < n; i++) {
        Px[i] = Py[i] = P[i];
    }
    int pointSize = sizeof(P[0]);
    qsort(Px, n, pointSize, Point::compareX);
    qsort(Py, n, pointSize, Point::compareY);
    return recursiveClosestPair(Px, Py, n);
}

Pair::Pair() {
    this->a = Point(0, 0);
    this->b = Point(0, 1);
    this->distance = Point::distance(a, b);
}

Pair::Pair(Point a, Point b) {
    this->a = a;
    this->b = b;
    this->distance = Point::distance(a, b);
}

Pair::Pair(Point a, Point b, double distance) {
    this->a = a;
    this->b = b;
    this->distance = distance;
}

Point Pair::getFirstPoint() const {
    return a;
}

Point Pair::getSecPoint() const {
    return b;
}

double Pair::getDistance() const {
    return distance;
}
