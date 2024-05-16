curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ' .[] | select( .id == 150) .name'
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ' .[] | select( .id == 150) .powerstats.power'
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ' .[] | select( .id == 150) .appearance.gender'