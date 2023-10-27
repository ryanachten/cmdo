import { h, render } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";

import App from "./components/App.js";

const html = htm.bind(h);

(function () {
  render(html`<${App} />`, document.body);
})();
