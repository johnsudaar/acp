class Redis < FPM::Cookery::Recipe
  homepage 'https://github.com/johnsudaar/acp'
  source   'https://github.com/johnsudaar/acp/archive/master.tar.gz'

  name     'acp'
  version  ENV['VERSION'] || 'dev'
  revision '1'

  description 'ACP Control Panel'

  config_files '/etc/acp/acp.yml'

  def build
    make 'VERSION' => ENV['VERSION'] || 'dev'
  end

  def install
    make :install, 'DESTDIR' => destdir
  end
end
