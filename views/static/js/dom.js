/**
 * @typedef {{
 *  container: HTMLElement,
 *  class: string
 * }} Command
 * @type {Object.<string, HTMLElement>}
 */
let commands = {};
const commandColors = ["cyan", "magenta", "green", "blue"];

/**
 * @param {string} commandName
 * @param {string} color
 * @returns {HTMLElement}
 */
const createCommandContainer = (commandName, color) => {
  const container = document.createElement("section");
  container.className = `container container--${color}`;

  const heading = document.createElement("span");
  heading.className = "container__heading";
  heading.innerText = commandName;

  container.appendChild(heading);
  container.appendChild(document.createElement("ul"));
  return container;
};

/**
 *
 * @param {import('./api.js').Message} message
 */
export const createMessageItem = (message) => {
  const main = document.getElementById("content");

  const content = message.messageBody.trim();
  if (!content) return;

  const item = document.createElement("li");
  item.innerText = content;
  if (message.messageType === "error") {
    item.classList.add("error");
  }

  if (!commands[message.commandName]) {
    const keyCount = Object.keys(commands).length;
    const color = commandColors[keyCount % commandColors.length];
    const container = createCommandContainer(message.commandName, color);
    commands[message.commandName] = container;
    main.appendChild(container);
  }

  const list = commands[message.commandName].querySelector("ul");
  list.appendChild(item);
  list.scroll({
    top: list.scrollHeight,
    behavior: "smooth",
  });
};
