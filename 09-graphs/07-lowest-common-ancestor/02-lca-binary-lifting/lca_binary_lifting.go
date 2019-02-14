package lca

//
// tin[v] : record in vertex time for v
// tout[v] : record out vertex time for v
// is_ancestor(u, v): tin[u] <= tin[v] && tout[u] >= tout[v]
// up[v][j] : record 2^j ancestor of v
//   up[v][0] = p
//   up[v][j] = up[ up[v][j-1] ][j-1]    // binary lifting
//
// 
// int lca(int u, int v)
// {
//     if (is_ancestor(u, v))
//         return u;
//     if (is_ancestor(v, u))
//         return v;
//     for (int i = l; i >= 0; --i) {
//         if (!is_ancestor(up[u][i], v))
//             u = up[u][i];
//     }
//     return up[u][0];
// }
// 
// 
