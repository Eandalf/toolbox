# To replace the command level env setting start command in CRA (PORT=8080 react-scripts start --watch)
# Set `start:dev:win: react-scripts start --watch` in package.json scripts section.
# To run Powershell scripts, run the following command in powershell with administrator:
# Get-ExecutionPolicy
# Set-ExecutionPolicy RemoteSigned
$env:PORT="8080"
npm run start:dev:win

# After the development, run the following command in powershell with administrator:
# Set-ExecutionPolicy Restricted
# Get-ExecutionPolicy

# Or, simply set PORT=8080 in .env file.
# CRA would read .env and set them.
# use `npm run start:dev:win` solely would also work.
