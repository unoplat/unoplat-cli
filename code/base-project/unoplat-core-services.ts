import { readFileSync } from 'fs';
import inquirer from 'inquirer';


 export default class UnoPlatCoreServices {   
    private coreServices: string[];

    constructor() {
        this.coreServices = JSON.parse(readFileSync('resources/core-services.json', 'utf-8'));
    }

    public async installCoreServices(): Promise<void> {
        const answers = await inquirer.prompt([
            {
                type: 'checkbox',
                name: 'services',
                message: 'Which core services would you like to install?',
                choices: this.coreServices,
            },
        ]);

        console.log(`You selected: ${answers.services.join(', ')}`);
    }
}
