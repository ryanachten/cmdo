import { useMemo } from "https://esm.sh/preact@10.18.1/hooks";

/**
 * @param {string} searchTerm
 * @param {import('./components/App.js').CommandHash} commands
 * @returns {import('./components/App.js').CommandHash}
 */
export const useFilteredCommands = (searchTerm, commands) =>
  useMemo(() => {
    if (!searchTerm) return commands;

    const term = searchTerm.toLowerCase();
    const updatedCommands = {};
    for (const key in commands) {
      updatedCommands[key] = {
        ...commands[key],
        history: commands[key].history.filter((cmd) =>
          cmd.messageBody.includes(term)
        ),
      };
    }
    return updatedCommands;
  }, [searchTerm, commands]);

/**
 * @param {string} searchTerm
 * @param {import('./components/App.js').Message[]} history
 * @returns {import('./components/App.js').Message[]}
 */
export const useFilteredHistory = (searchTerm, history) =>
  useMemo(() => {
    if (!searchTerm) return history;

    const term = searchTerm.toLowerCase();
    return history.filter((cmd) =>
      cmd.messageBody.toLowerCase().includes(term)
    );
  }, [searchTerm, history]);

/**
 * Removes all ASCII escape strings from terminal output
 * @param {string} rawString
 * @returns {string}
 */
export const formatMessageBody = (rawString) =>
  rawString.replace(/\u001b\[\d+m/g, "").trim();
