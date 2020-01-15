#include <functional>
#include <mutex>
#include <thread>
#include <iostream>
#include <chrono>

using namespace std;
class FooBar {
private:
    int n;
    bool turnFoo = true;
    mutex m;
    condition_variable cv;

public:
    FooBar(int n) {
        this->n = n;
    }

    void foo(function<void()> printFoo) {
        unique_lock<mutex> lk(m);
        for (int i = 0; i < n; i++) {
        	// printFoo() outputs "foo". Do not change or remove this line.
            cv.wait(lk, [this] { return this->turnFoo; });
        	printFoo();
            turnFoo = false;
            cv.notify_one();
        }
    }

    void bar(function<void()> printBar) {
        unique_lock<mutex> lk(m);
        for (int i = 0; i < n; i++) {
        	// printBar() outputs "bar". Do not change or remove this line.
            cv.wait(lk, [this] { return !this->turnFoo; });
        	printBar();
            turnFoo = true;
            cv.notify_one();
        }
    }
};

void pf() {
    // cout << "foo";
}
void pb() {
    // cout << "bar";
}

int main() {
    auto start = chrono::high_resolution_clock::now();
    FooBar *foobar = new FooBar(1000);
    auto t1 = thread([foobar] {
        foobar->bar(pb);
    });
    auto t2 = thread([foobar] {
        foobar->foo(pf);
    });
    t1.join();
    t2.join();
    auto elapsed = chrono::high_resolution_clock::now() - start;
    cout << "wait for "
        << chrono::duration_cast<chrono::microseconds>(elapsed).count()
        << " ms\n";
    return 0;
}