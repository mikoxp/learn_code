require 'pg'

begin

    con = PG.connect :dbname => 'postgres', :user => 'postgres' ,:password => 'postgres'

    rs = con.exec 'SELECT VERSION()'
    puts rs.getvalue 0, 0
    
rescue PG::Error => e

    puts e.message 
    
ensure

    con.close if con
    
end