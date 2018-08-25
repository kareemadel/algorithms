//
// Created by kareemadel on 8/20/18.
//

#include <cmath>
#include <Point.h>

Point::Point() {
    this->x = 0.0;
    this->y = 0.0;
}

Point::Point(double x, double y) {
    this->x = x;
    this->y = y;
}

double Point::distance(const Point &p1, const Point &p2) {
    return sqrt((p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y));
}

double Point::distance() const {
    return Point::distance(*this, Point());
}

double Point::distance(const Point &p) const {
    return Point::distance(*this, p);
}

int Point::compareX(const void* a, const void* b) {
    Point *p1 = (Point *) a, *p2 = (Point *) b;
    if (p1->x > p2->x) {
        return 1;
    } else if (p1->x < p2->x) {
        return -1;
    } else {
        return 0;
    }
}

int Point::compareY(const void* a, const void* b) {
    Point *p1 = (Point *) a, *p2 = (Point *) b;
    if (p1->y > p2->y) {
        return 1;
    } else if (p1->y < p2->y) {
        return -1;
    } else {
        return 0;
    }
}

int Point::compareX() const {
    Point p = Point();
    return Point::compareX(this, &p);
}

int Point::compareX(const Point &p) const {
    return Point::compareX(this, &p);
}

int Point::compareY() const {
    Point p = Point();
    return Point::compareY(this, &p);
}

int Point::compareY(const Point &p) const {
    return Point::compareY(this, &p);
}

double Point::getX() const {
    return this->x;
}

double Point::getY() const {
    return this->y;
}
