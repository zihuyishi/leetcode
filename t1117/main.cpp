#include <functional>
#include <thread>
#include <iostream>
#include <mutex>
#include <vector>

using namespace std;

class H2O {
private:
    mutex m;
    condition_variable cv;
    int state;
public:
    H2O() {
        state = 0;
    }

    void hydrogen(function<void()> releaseHydrogen) {
        unique_lock<mutex> lk(m);
        while (state == -2) {
            // cout << "wait h\n";
            cv.wait(lk);
        }
        // releaseHydrogen() outputs "H". Do not change or remove this line.
        releaseHydrogen();
        state -= 1;
        cv.notify_all();
    }

    void oxygen(function<void()> releaseOxygen) {
        unique_lock<mutex> lk(m);

        while (state > 0) {
            // cout << "wait o\n";
            cv.wait(lk);
        } 
        // releaseOxygen() outputs "O". Do not change or remove this line.
        releaseOxygen();
        state += 2;
        cv.notify_all();
    }
};

void pH() {
    cout << 'h' << endl;
}
void pO() {
    cout << 'o' << endl;
}

void callIt(H2O *h, char c) {
    if (c == 'h') {
        // cout << "push h\n";
        h->hydrogen(pH);
    } else {
        // cout << "push o\n";
        h->oxygen(pO);
    }
}

void processVec(H2O *h, vector<char>& s) {
    vector<thread*> ts;
    for (auto c : s) {
        auto t = new thread([h, c] {
            callIt(h, c);
        });
        ts.push_back(t);
    }
    for (thread *t : ts) {
        t->join();
    }
}

int main() {
    H2O *h = new H2O();
    auto t1 = thread([h] {
        auto s = vector<char> {
            'h', 'h', 'h', 'o', 'o',
        };
        processVec(h, s);
    });
    auto t2 = thread([h] {
        auto s = vector<char> {
            'h', 'h', 'h', 'o',
        };
        processVec(h, s);
    });
    t1.join();
    t2.join();
    delete h;
    return 0;
}