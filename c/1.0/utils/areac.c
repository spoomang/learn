
/*
    Generating library for area.
     - gcc -c *.c
     - ar -cvq libxutils.a *.o

*/

double getTraingleArea(double height, double base){
    return 0.5 * height * base;
}

double getCircleArea(double radius) {
    return 3.14 * radius * radius;
}