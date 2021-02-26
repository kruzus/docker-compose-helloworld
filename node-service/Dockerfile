#Docker files must begin with FROM keyword, in order to do its magic.
FROM node:15.10.0-alpine3.11

# Established working directory; "app" you can name it anything you want: apple, potato, etc..
WORKDIR /app

# Now since we created a base directory, we can copy from the folder we are in right now
# To our contrainer into the /app folder.

# Copy takes 2 arguments think it as: COPY FROM TO

#COPY CONTENTS FROM WHERE WE ARE(in this case node-docer/) and file that has "package" and ends
# with ".json" format, THEN we can copy that to our container directory
# Think it as, copy package.json to ./app/{files rest here}
COPY package*.json ./

#Then we copy finished prodcut 
COPY ./dist ./dist

# Running NPM I to install dependecny (expressjs)
RUN npm i

#and simple running our project.
CMD ["node", "./dist/server.js"]


EXPOSE 8080
