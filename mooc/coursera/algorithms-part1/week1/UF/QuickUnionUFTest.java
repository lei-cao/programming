import static org.junit.jupiter.api.Assertions.*;

class QuickUnionUFTest {

    @org.junit.jupiter.api.Test
    void connected() {
        QuickUnionUF uf = new QuickUnionUF(10);
        uf.union(2, 9);
        uf.union(4,9);
        uf.union(3,4);
        uf.union(6,5);

        assertTrue(uf.connected(3,9));
        assertTrue(uf.connected(5,6));
        assertTrue(uf.connected(2,3));
        assertTrue(uf.connected(2,4));
        assertFalse(uf.connected(8,9));
        assertFalse(uf.connected(8,5));
        assertFalse(uf.connected(8,1));
    }
}