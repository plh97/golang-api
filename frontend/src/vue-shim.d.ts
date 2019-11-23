import Vue from 'vue';
import { AxiosStatic } from 'axios';

// vue-shim.d.ts
declare module 'vue/types/vue' {
  interface Vue {
    $http: AxiosStatic;
    validate: () => boolean
  }
}
