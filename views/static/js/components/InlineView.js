import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";

const html = htm.bind(h);

/**
 * @typedef {{commandName: string, messageBody: string, messageType: "error" | "information"}} Message
 * @param {{ history: import('./App.js').Message[] }}
 */
function InlineView({ history }) {
  return html`<ul>
    ${history.map(
      ({ commandName, messageBody }) =>
        html`<li><span>${commandName}</span> ${messageBody}</li>`
    )}
  </ul>`;
}

export default InlineView;
