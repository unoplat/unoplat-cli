import { readFileSync } from 'fs';
import inquirer from 'inquirer';
import UnoPlatCoreServices from './unoplat-core-services.js'; // Update the path to the actual location of the CoreServices file



export default class UnoplatCLI {
    private actions: string[];
    private unoplatCoreServices: UnoPlatCoreServices;

    constructor() {
        this.actions = JSON.parse(readFileSync('resources/actions.json', 'utf-8'));
        this.unoplatCoreServices = new UnoPlatCoreServices();
    }

    async showActions() {
        const answers = await inquirer.prompt([
            {
                type: 'list',
                name: 'action',
                message: 'What action would you like to perform?',
                choices: this.actions,
            },
        ]);

        console.log(`You selected: ${answers.action}`);

        if (answers.action == "install-core-services"){
            await this.unoplatCoreServices.installCoreServices();
        }
    }
}
