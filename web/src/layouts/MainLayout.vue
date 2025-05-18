<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Quasar App
        </q-toolbar-title>

        <div>Quasar v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item-label
          header
        >
          Essential Links
        </q-item-label>

        <MenuItem
          v-for="mil in menuItemlist"
          :key="mil.title"
          v-bind="mil"
        />

        <!-- <EssentialLink
          v-for="link in linksList"
          :key="link.title"
          v-bind="link"
        /> -->
      </q-list>
    </q-drawer>

    <!-- <q-drawer
        side="right"
        v-model="rightDrawer.show"
        show-if-above
        bordered
        overlay
        :width="$q.screen.width / 3 - 1"
        :breakpoint="500"
        behavior="mobile"
        no-swipe-open
        no-swipe-close
        no-swipe-backdrop
        :class="$q.dark.isActive ? 'bg-grey-9' : 'bg-grey-3'"
        @hide="rightDrawer.toggle(false)"
      >
        <q-scroll-area class="fit">
          <div class="q-pa-sm">
            <div v-for="n in 20" :key="n">Drawer {{ n }} / 50</div>
          </div>
        </q-scroll-area>
      </q-drawer> -->

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import MenuItem, { type MenuItemProps } from 'components/MenuItem.vue';
// import { useRightDrawerStore } from 'src/stores/right-drawer-store';

// const rightDrawer = useRightDrawerStore()

const menuItemlist: MenuItemProps[] = [
  {
    title: 'Home',
    to: '/',
    icon: 'school'
  },
  {
    title: 'Foods',
    to: '/foods',
    icon: 'school'
  }
]

const leftDrawerOpen = ref(false);

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
