#include <iostream>

using namespace std;
class Bouncingball
{
public:
    static int bouncingBall(double h, double bounce, double window) {
        if (h > 0 && bounce > 0 && bounce < 1 && window < h) {
            int total_distance = 1;  // Ball falls initially
            while (h * bounce > window) {
                h *= bounce;  // Ball bounces to two-thirds of previous height
                total_distance += 2;  // Count the ball going up and down
            }
            return total_distance;
        } else {
            return -1;  // Invalid input
        }
    }
};

int main (){
    double h,bounce,window;
    cin>>h>>bounce>>window;
    cout<<Bouncingball::bouncingBall(h,bounce,window);
}