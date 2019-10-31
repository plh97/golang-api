import _Vue from "vue";
import Axios from "axios";

// export function AxiosPlugin<AxiosPlugOptions>(
//   Vue: typeof _Vue,
//   options?: AxiosPluginOptions
// ): void {
//   Vue.prototype.$http = Axios;
// }

export class AxiosPluginOptions {
  // add stuff
}

import { AxiosStatic } from "axios";

declare module "vue/types/vue" {
  interface Vue {
    $http: AxiosStatic;
  }
}