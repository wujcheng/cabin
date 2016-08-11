/tmp/supervisor_install.sh:
  file.managed:
    - source: salt://install/source/scripts/supervisor_install.sh
    - user: root
    - group: root
    - mode: 755
  cmd.run:
    - user: root
    - shell: /bin/bash