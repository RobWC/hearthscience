API Overview
============

Calls
=====

/v1/<CALL>
----------

All calls must be prefaced with a version id for non-HTML outputs

Query options
-------------

-	type
	1.	json
		-	Returns the call in json format
	2.	xml
		-	Returns the call in xml format
-	key
	-	Provides a key for api authentication

GET /cards
----------

Returns the complete list of cards

GET /card/<id>
--------------

Returns a card based upon ID

GET /minions
------------

Returns the complete list of minions

GET /minion/<ID>
----------------

Returns a minion by ID

GET /spells
-----------

Returns the complete list of spells

GET /spell/<ID>
---------------

Returns a spell by ID

GET /enchantments
-----------------

Returns the complete list of enchantments

GET /enchantment/<ID>
---------------------

Returns an enchantment by ID

GET /heros
----------

Returns the complete list of heros

GET /hero/<ID>
--------------

Returns a Hero by ID

GET /heropowers
---------------

Returns the complete list of hero powers

GET /heropower/<id>
-------------------

Returns a Hero Power by ID

GET /weapons
------------

Returns the complete list of weapons

GET /weapon/<id>
----------------

Returns a weapon by ID

GET /rarity/<rarity>
--------------------

Returns all cards matching specific rarity

GET /cost/<cost>
----------------

Returns all cards matching specific cost

GET /attack/<attack>
--------------------

Returns all cards matching specific attack

GET /health/<health>
--------------------

Returns all cards matching specific health

GET /mechanic/<mechanic>
------------------------

Returns all cards matching specific mechanic

GET /race/<race>
----------------

Returns all cards matching specific race
