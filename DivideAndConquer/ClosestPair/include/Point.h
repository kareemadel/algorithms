//
// Created by kareemadel on 8/20/18.
//

#ifndef CLOSESTPAIR_POINT_H
#define CLOSESTPAIR_POINT_H

#endif //CLOSESTPAIR_POINT_H

class Point {
    double x, y;
public:
    Point();
    Point(double x, double y);

    double distance() const;
    double distance(const Point &p) const;
    int compareX() const;
    int compareX(const Point &p) const;
    int compareY() const;
    int compareY(const Point &p) const;
    double getX() const;
    double getY() const;
    static double distance(const Point &p1, const Point &p2);
    static int compareX(const void* a, const void* b);
    static int compareY(const void* a, const void* b);
};
