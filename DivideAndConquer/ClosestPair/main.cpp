#include <iostream>
#include <Point.h>
#include <Pair.h>
using namespace std;

std::ostream &operator<<(std::ostream &out, const Point &p) {
    out << '(' << p.getX() << ',';
    out << p.getY() << ')';
    return out;
}

std::ostream &operator<<(std::ostream &out, const Pair &p) {
    out << "First Point: " << p.getFirstPoint() << endl;
    out << "Second Point: " << p.getSecPoint() << endl;
    out << "Distance: " << p.getDistance() << endl;
    return out;
}

int main() {
    Point P[] = {{2,  3},
                 {12, 3},
                 {40, 50},
                 {5,  3},
                 {12, 10},
                 {12, 11},
                 {3,  4},
                 {10, 20}};
    int size = sizeof(P) / sizeof(P[0]);
    Pair closestPair = Pair::findClosestPair(P, size);
    cout << closestPair << endl;
    return 0;
}