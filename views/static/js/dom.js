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

  const heading = document.createElement("h2");
  heading.innerText = commandName;

  container.appendChild(heading);
  container.appendChild(document.createElement("ul"));
  return container;
};

/**
 *
 * @param {Message} message
 */
export const createMessageItem = (message) => {
  const main = document.getElementById("content");

  const item = document.createElement("li");
  item.innerText = message.messageBody;

  if (!commandContainers[message.commandName]) {
    const container = createCommandContainer(message.commandName);
    commandContainers[message.commandName] = container;
    main.appendChild(container);
  }
  commandContainers[message.commandName].querySelector("ul").appendChild(item);
};
