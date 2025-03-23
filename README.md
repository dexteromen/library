Here are the steps to update Node.js using nvm:

1. Install nvm (if not already installed):

```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash
```

2. Load nvm (you may need to restart your terminal or run the following command):

```
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
```

3. Install Node.js version 20.0.0:

```
nvm install 20.0.0
```

4. Use Node.js version 20.0.0:

```
nvm use 20.0.0
```

5. Verify the Node.js version:

```
node -v
```

After updating Node.js, you can try running your npm commands again.

# Checking Outdated package

Check for outdated packages:

```
npm outdated
```

Update all dependencies to the latest major versions:

```
npx npm-check-updates -u
npm install
```
