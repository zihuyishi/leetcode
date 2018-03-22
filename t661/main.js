class Counter {
    constructor() {
        this.sum = 0;
        this.count = 0;
    }

    add(v) {
        this.sum += v;
        this.count += 1;
    }

    value() {
        return Math.floor(this.sum / this.count);
    }
}

const dir = [[-1,-1],[-1,0],[-1,1],
    [0,-1],[0,0],[0,1],
    [1,-1],[1,0],[1,1]];

/**
 * @param {number[][]} M
 * @return {number[][]}
 */
const imageSmoother = function(M) {
    const A = M.map(l => {
        return l.map(() => 0);
});
    for (let i = 0; i < M.length; i++) {
        for (let j = 0; j < M[i].length; j++) {
            const counter = new Counter();
            for (const d of dir) {
                const x = i + d[0];
                const y = j + d[1];
                if (x >= 0 && y >= 0 && x < M.length && y < M[i].length) {
                    counter.add(M[x][y]);
                }
            }
            A[i][j] = counter.value();
        }
    }
    return A;
};