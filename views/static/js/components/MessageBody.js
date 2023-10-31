import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import Convert from "https://esm.sh/ansi-to-html@0.7.2";

const html = htm.bind(h);
const convert = new Convert();

/**
 * @param {{messageBody: string}}
 */
function MessageBody({ messageBody }) {
  // TODO: either avoid dangerouslySetInnerHTML or sanitise the output
  const formattedMessage = convert.toHtml(messageBody);
  return html`<span
    dangerouslySetInnerHTML="${{ __html: formattedMessage }}"
  />`;
}

export default MessageBody;
