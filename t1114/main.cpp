#include <functional>
#include <mutex>
#include <thread>
#include <iostream>
using namespace std;

class Foo {
    mutex lock1;
    mutex lock2;
public:
    Foo() {
        std::lock(lock1, lock2);
    }

    void first(function<void()> printFirst) {
        
        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        lock1.unlock();
    }

    void second(function<void()> printSecond) {
        lock1.lock();
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        lock2.unlock(); 
    }

    void third(function<void()> printThird) {
        lock2.lock(); 
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
    }
};

int main() {
    Foo *f = new Foo();
    auto t1 = thread([f] {
        f->third([] {
            cout << 3 << endl;            
        });
    });
    auto t2 = thread([f] {
        f->second([] {
            cout << 2 << endl;
        });
    });
    auto t3 = thread([f] {
        f->first([] {
            this_thread::sleep_for(2s);
            cout << 1 << endl;
        });
    });
    t1.join();
    t2.join();
    t3.join();
    return 0;
}