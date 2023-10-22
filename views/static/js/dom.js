/**
 * @type {Object.<string, HTMLElement>}
 */
let commandContainers = {};

/**
 * @param {string} commandName
 * @returns {HTMLElement}
 */
const createCommandContainer = (commandName) => {
  const container = document.createElement("section");
  container.className = "container";

  const heading = document.createElement("span");
  heading.className = "container__heading";
  heading.innerText = "⚡ " + commandName;

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

  const item = document.createElement("li");
  item.innerText = message.messageBody;
  if (message.messageType === "error") {
    item.classList.add("error");
  }

  if (!commandContainers[message.commandName]) {
    const container = createCommandContainer(message.commandName);
    commandContainers[message.commandName] = container;
    main.appendChild(container);
  }

  const list = commandContainers[message.commandName].querySelector("ul");
  list.appendChild(item);
  list.scroll({
    top: list.scrollHeight,
    behavior: "smooth",
  });
};
