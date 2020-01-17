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

public:
    FooBar(int n) {
        this->n = n;
    }

    void foo(function<void()> printFoo) {
        for (int i = 0; i < n; i++) {
            while (!turnFoo)
                this_thread::yield();
        	// printFoo() outputs "foo". Do not change or remove this line.
        	printFoo();
            turnFoo = false;
        }
    }

    void bar(function<void()> printBar) {
        for (int i = 0; i < n; i++) {
            while (turnFoo)
                this_thread::yield();
        	// printBar() outputs "bar". Do not change or remove this line.
        	printBar();
            turnFoo = true;
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