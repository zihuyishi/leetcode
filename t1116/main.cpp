#include <functional>
#include <iostream>
#include <mutex>
#include <thread>
using namespace std;

class ZeroEvenOdd
{
private:
    int cur;
    int n;
    mutex m;
    condition_variable cv;
    bool turnZero = true;

public:
    ZeroEvenOdd(int n)
    {
        this->n = n;
        this->cur = 1;
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber)
    {
        unique_lock<mutex> lk(m);
        while (cur <= n) {
            cv.wait(lk, [this] {
                return this->turnZero;
            });
            if (cur <= n) {
                printNumber(0);
            }
            this->turnZero = false;
            cv.notify_all();
        }
    }

    void even(function<void(int)> printNumber)
    {
        unique_lock<mutex> lk(m);
        cv.notify_all();
        while (cur <= n) {
            cv.wait(lk, [this] {
                return (!this->turnZero) && ((this->cur % 2) == 0);
            });
            if (cur <= n) {
                printNumber(cur);
            }
            cur += 1;
            this->turnZero = true;
            cv.notify_all();
        }
    }

    void odd(function<void(int)> printNumber)
    {
        unique_lock<mutex> lk(m);
        while (cur <= n) {
            cv.wait(lk, [this] {
                return (!this->turnZero) && ((this->cur % 2) == 1);
            });
            if (cur <= n) {
                printNumber(cur);
            }
            cur += 1;
            this->turnZero = true;
            cv.notify_all();
        }
    }
};

void pn(int i) {
    cout << i;
}

static const int N = 0;
int main() {
    ZeroEvenOdd *p = new ZeroEvenOdd(N);
    auto odd = thread([p] {
        for (int i = 0; i < (N+1)/2; i++) {
            p->odd(pn);
        }
    });
    auto even = thread([p] {
        for (int i = 0; i < N / 2; i++) {
            p->even(pn);
        };
    });
    auto zero = thread([p] {
        for (int i = 0; i < N; i++) {
            p->zero(pn); 
        }
    });
    zero.join();
    odd.join();
    even.join();
    delete p;
    return 0;
}