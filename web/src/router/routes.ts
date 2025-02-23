import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'foods', component: () => import('pages/FoodsPage.vue') },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

// const route = (path: string, component: string, children: { path: string, component: string }[] = []): RouteRecordRaw => ({
//   path,
//   component: () => import(`../${component}`),
//   children: children.map(child => ({
//     path: child.path,
//     component: () => import(`../${child.component}`),
//   })),
// });

// const routes: RouteRecordRaw[] = [
//   route('/', 'layouts/MainLayout.vue', [{ path: '', component: 'pages/IndexPage.vue' }]),
//   route('/:catchAll(.*)*', 'pages/ErrorNotFound.vue'),
// ];

export default routes;
