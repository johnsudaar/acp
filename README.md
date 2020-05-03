# ACP

## Installation

1. Create a file in `/etc/apt/sources.list.d/` named `acp.list` and add the following content:

```
deb https://packages.hurter.fr/stable/ buster main
```

2. Import the GPG Key:

```
sudo apt-key adv --keyserver pgp.mit.edu --recv-keys 740600ADC672D1B3
```

3. Update your repository

```
apt update
```

4. Install ACP

```
apt install acp
```

5. Install MongoDB: [link](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)
