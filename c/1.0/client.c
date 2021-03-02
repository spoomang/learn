
#include <stdio.h>
#include "headers/area.h"

int main() {
    printf("programme starts\n");

    double areaOfCircle = getCircleArea(4);
    double areaOfTraingle = getTraingleArea(5, 10);

    printf("Area of circle = %f\n", areaOfCircle);
    printf("Area of traingle = %f\n", areaOfTraingle);

    return 0;
}