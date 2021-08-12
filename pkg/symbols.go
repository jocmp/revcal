package revcal

var Symbols = []string{
	"Raisin (Grape)",
	"Safran (Saffron)",
	"Châtaigne (Chestnut)",
	"Colchique (Crocus)",
	"Cheval (Horse)",
	"Balsamine (Impatiens)",
	"Carotte (Carrot)",
	"Amarante (Amaranth)",
	"Panais (Parsnip)",
	"Cuve (Vat)",
	"Pomme de terre (Potato)",
	"Immortelle (Strawflower)",
	"Potiron (Calabaza)",
	"Réséda (Mignonette)",
	"Âne (Donkey)",
	"Belle de nuit (The four o’clock flower)",
	"Citrouille (Pumpkin)",
	"Sarrasin (Buckwheat)",
	"Tournesol (Sunflower)",
	"Pressoir (Wine-Press)",
	"Chanvre (Hemp)",
	"Pêche (Peach)",
	"Navet (Turnip)",
	"Amaryllis (Amaryllis)",
	"Bœuf (Cow)",
	"Aubergine (Eggplant)",
	"Piment (Chili Pepper)",
	"Tomate (Tomato)",
	"Orge (Barley)",
	"Tonneau (Barrel)",
	"Pomme (Apple)",
	"Céleri (Celery)",
	"Poire (Pear)",
	"Betterave (Beet root)",
	"Oie (Goose)",
	"Héliotrope (Heliotrope)",
	"Figue (Fig)",
	"Scorsonère (Black Salsify)",
	"Alisier (Chequer Tree)",
	"Charrue (Plough)",
	"Salsifis (Salsify)",
	"Macre (Water chestnut)",
	"Topinambour (Jerusalem Artichoke)",
	"Endive (Endive)",
	"Dindon (Turkey)",
	"Chervis (Skirret)",
	"Cresson (Watercress)",
	"Dentelaire (Leadworts)",
	"Grenade (Pomegranate)",
	"Herse (Harrow)",
	"Bacchante (Asarum baccharis)",
	"Azerole (Acerola)",
	"Garance (Madder)",
	"Orange (Orange)",
	"Faisan (Pheasant)",
	"Pistache (Pistachio)",
	"Macjonc (Tuberous pea)",
	"Coing (Quince)",
	"Cormier (Service tree)",
	"Rouleau (Roller)",
	"Raiponce (Rampion)",
	"Turneps (Turnip)",
	"Chicorée (Chicory)",
	"Nèfle (Medlar)",
	"Cochon (Pig)",
	"Mâche (Corn Salad)",
	"Chou-fleur (Cauliflower)",
	"Miel (Honey)",
	"Genièvre (Juniper)",
	"Pioche (Pickaxe)",
	"Cire (Wax)",
	"Raifort (Horseradish)",
	"Cèdre (Cedar tree)",
	"Sapin (Fir tree)",
	"Chevreuil (Roe Deer)",
	"Ajonc (Gorse)",
	"Cyprès (Cypress Tree)",
	"Lierre (Ivy)",
	"Sabine (Juniper)",
	"Hoyau (Grub-hoe)",
	"Érable sucré (Maple Tree)",
	"Bruyère (Heather)",
	"Roseau (Reed plant)",
	"Oseille (Sorrel)",
	"Grillon (Cricket)",
	"Pignon (Pinenut)",
	"Liège (cork)",
	"Truffe (Truffle)",
	"Olive (Olive)",
	"Pelle (shovel)",
	"Tourbe (Peat)",
	"Houille (Coal)",
	"Bitume (Bitumen)",
	"Soufre (Sulphur)",
	"Chien (Dog)",
	"Lave (Lava)",
	"Terre végétale (Topsoil)",
	"Fumier (Manure)",
	"Salpêtre (Saltpeter)",
	"Fléau (Flail)",
	"Granit (Granite stone)",
	"Argile (Clay)",
	"Ardoise (Slate)",
	"Grès (Sandstone)",
	"Lapin (Rabbit)",
	"Silex (Flint)",
	"Marne (Marl)",
	"Pierre à chaux (Limestone)",
	"Marbre (Marble)",
	"Van (Winnowing basket)",
	"Pierre à plâtre (Gypsum)",
	"Sel (Salt)",
	"Fer (Iron)",
	"Cuivre (Copper)",
	"Chat (Cat)",
	"Étain (Tin)",
	"Plomb (Lead)",
	"Zinc (Zinc)",
	"Mercure (Mercury (metal))",
	"Crible (Sieve)",
	"Lauréole (Spurge-laurel)",
	"Mousse (Moss)",
	"Fragon (Butcher’s Broom)",
	"Perce-neige (Snowdrop)",
	"Taureau (Bull)",
	"Laurier-thym (Laurustinus)",
	"Amadouvier (Tinder polypore)",
	"Mézéréon (Daphne mezereum)",
	"Peuplier (Poplar Tree)",
	"Coignée (Axe)",
	"Ellébore (Hellebore)",
	"Brocoli (Broccoli)",
	"Laurier (Laurel)",
	"Avelinier (Cob or filbert)",
	"Vache (Cow)",
	"Buis (Box Tree)",
	"Lichen (Lichen)",
	"If (Yew tree)",
	"Pulmonaire (Lungwort)",
	"Serpette (Billhook)",
	"Thlaspi (Pennycress)",
	"Thimelé (Rose Daphne)",
	"Chiendent (Couch Grass)",
	"Trainasse (Knotweed)",
	"Lièvre (Hare)",
	"Guède (Woad)",
	"Noisetier (Hazel)",
	"Cyclamen (Cyclamen)",
	"Chélidoine (Celandine)",
	"Traîneau (Sleigh)",
	"Tussilage (Coltsfoot)",
	"Cornouiller (Dogwood)",
	"Violier (Matthiola)",
	"Troène (Privet)",
	"Bouc (Billygoat)",
	"Asaret (Wild Ginger)",
	"Alaterne (Buckthorn)",
	"Violette (Violet (plant))",
	"Marceau (Goat Willow)",
	"Bêche (Spade)",
	"Narcisse (Narcissus)",
	"Orme (Elm Tree)",
	"Fumeterre (Common fumitory)",
	"Vélar (Hedge Mustard)",
	"Chèvre (Goat)",
	"Épinard (Spinach)",
	"Doronic (Large-flowered Leopard’s Bane)",
	"Mouron (Pimpernel)",
	"Cerfeuil (Chervil)",
	"Cordeau (Twine)",
	"Mandragore (Mandrake)",
	"Persil (Parsley)",
	"Cochléaria (Scurvy-grass)",
	"Pâquerette (Daisy)",
	"Thon (Tuna)",
	"Pissenlit (Dandelion)",
	"Sylve (Forest)",
	"Capillaire (Maidenhair fern)",
	"Frêne (Ash Tree)",
	"Plantoir (Dibber: a hand gardening tool)",
	"Primevère (Primrose)",
	"Platane (Plane Tree)",
	"Asperge (Asparagus)",
	"Tulipe (Tulip)",
	"Poule (Hen)",
	"Bette (Chard Plant)",
	"Bouleau (Birch Tree)",
	"Jonquille (Daffodil)",
	"Aulne (Alder)",
	"Couvoir (Hatchery)",
	"Pervenche (Periwinkle)",
	"Charme (Ironwood)",
	"Morille (Morel)",
	"Hêtre (Beech Tree)",
	"Abeille (Bee)",
	"Laitue (Lettuce)",
	"Mélèze (Larch)",
	"Ciguë (Hemlock)",
	"Radis (Radish)",
	"Ruche (Hive)",
	"Gainier (Judas tree)",
	"Romaine (Lettuce)",
	"Marronnier (Chestnut Oak)",
	"Roquette (Arugula or Rocket)",
	"Pigeon (Pigeon)",
	"Lilas (Lilac)",
	"Anémone (Anemone)",
	"Pensée (Pansy)",
	"Myrtille (Blueberry)",
	"Greffoir (Knife)",
	"Rose (Rose)",
	"Chêne (Oak Tree)",
	"Fougère (Fern)",
	"Aubépine (Hawthorn)",
	"Rossignol (Nightingale)",
	"Ancolie (Columbine)",
	"Muguet (Lily of the Valley)",
	"Champignon (Button mushroom)",
	"Hyacinthe (Hyacinth)",
	"Râteau (Rake)",
	"Rhubarbe (Rhubarb)",
	"Sainfoin (Sainfoin)",
	"Bâton-d’or (Wallflower)",
	"Chamérops (Palm tree)",
	"Ver à soie (Silkworm)",
	"Consoude (Comfrey)",
	"Pimprenelle (Salad Burnet)",
	"Corbeille d’or (Basket of Gold)",
	"Arroche (Orache)",
	"Sarcloir (Garden hoe)",
	"Statice (Sea Lavender)",
	"Fritillaire (Fritillary)",
	"Bourrache (Borage)",
	"Valériane (Valerian)",
	"Carpe (Carp)",
	"Fusain (Spindle (shrub))",
	"Civette (Chive)",
	"Buglosse (Bugloss)",
	"Sénevé (Wild mustard)",
	"Houlette (Shepherd’s crook)",
	"Luzerne (Alfalfa)",
	"Hémérocalle (Daylily)",
	"Trèfle (Clover)",
	"Angélique (Angelica)",
	"Canard (Duck)",
	"Mélisse (Lemon Balm)",
	"Fromental (Oat grass)",
	"Martagon (Martagon lily)",
	"Serpolet (Thyme plant)",
	"Faux (Scythe)",
	"Fraise (Strawberry)",
	"Bétoine (Woundwort)",
	"Pois (Pea)",
	"Acacia (Acacia)",
	"Caille (Quail)",
	"Œillet (Carnation)",
	"Sureau (Elderberry)",
	"Pavot (Poppy plant)",
	"Tilleul (Linden or Lime tree)",
	"Fourche (Pitchfork)",
	"Barbeau (Cornflower)",
	"Camomille (Camomile)",
	"Chèvrefeuille (Honeysuckle)",
	"caille-lait (Bedstraw)",
	"Tanche (Tench)",
	"Jasmin (Jasmine Plant)",
	"Verveine (Verbena)",
	"Thym (Thyme Plant)",
	"Pivoine (Peony Plant)",
	"Chariot (Hand Cart)",
	"Seigle (Rye)",
	"Avoine (Oats)",
	"Oignon (Onion)",
	"Véronique (Speedwell)",
	"Mulet (Mule)",
	"Romarin (Rosemary)",
	"Concombre (Cucumber)",
	"Échalote (Shallot)",
	"Absinthe (Wormwood)",
	"Faucille (Sickle)",
	"Coriandre (Coriander)",
	"Artichaut (Artichoke)",
	"Girofle (Clove)",
	"Lavande (Lavender)",
	"Chamois (Chamois)",
	"Tabac (Tobacco)",
	"Groseille (Currant)",
	"Gesse (Hairy Vetchling)",
	"Cerise (Cherry)",
	"Parc (Park)",
	"Menthe (Mint)",
	"Cumin (Cumin)",
	"Haricot (Bean)",
	"Orcanète (Alkanet)",
	"Pintade (Guinea fowl)",
	"Sauge (Sage Plant)",
	"Ail (Garlic)",
	"Vesce (Tare)",
	"Blé (Wheat)",
	"Chalémie (Shawm)",
	"Épeautre (Einkorn Wheat)",
	"Bouillon blanc (Common Mullein)",
	"Melon (Honeydew Melon)",
	"Ivraie (Ryegrass)",
	"Bélier (Ram)",
	"Prêle (Horsetail)",
	"Armoise (Mugwort)",
	"Carthame (Safflower)",
	"Mûre (Blackberry)",
	"Arrosoir (Watering Can)",
	"Panis (Panic grass)",
	"Salicorne (Common Glasswort)",
	"Abricot (Apricot)",
	"Basilic (Basil)",
	"Brebis (Ewe)",
	"Guimauve (Marshmallow root)",
	"Lin (Flax)",
	"Amande (Almond)",
	"Gentiane (Gentian)",
	"Écluse (Lock)",
	"Carline (Carline thistle)",
	"Câprier (Caper)",
	"Lentille (Lentil)",
	"Aunée (Yellow starwort)",
	"Loutre (Otter)",
	"Myrte (Myrtle)",
	"Colza (Rapeseed)",
	"Lupin (Lupin)",
	"Coton (Cotton)",
	"Moulin (Mill)",
	"Prune (Plum)",
	"Millet (Millet)",
	"Lycoperdon (Puffball)",
	"Escourgeon (Six-row Barley)",
	"Saumon (Salmon)",
	"Tubéreuse (Tuberose)",
	"Sucrion (Sugar melon)",
	"Apocyn (Apocynum)",
	"Réglisse (Liquorice)",
	"Échelle (Ladder)",
	"Pastèque (Watermelon)",
	"Fenouil (Fennel)",
	"Épine vinette (Barberry)",
	"Noix (Walnut)",
	"Truite (Trout)",
	"Citron (Lemon)",
	"Cardère (Teasel)",
	"Nerprun (Buckthorn)",
	"Tagette (Mexican Marigold)",
	"Hotte (Sack)",
	"Églantine (Wild Rose)",
	"Noisette (Hazelnut)",
	"Houblon (Hops)",
	"Sorgho (Sorghum)",
	"Écrevisse (Crayfish)",
	"Bigarade (Bitter Orange)",
	"Verge d’or (Goldenrod)",
	"Maïs (Maize or Corn)",
	"Marron (Chestnut)",
	"Panier (Basket)",
	"La Fête de la Vertu (Celebration of Virtue)",
	"La Fête du Génie (Celebration of Talent)",
	"La Fête du Travail (Celebration of Labour)",
	"La Fête de l'Opinion (Celebration of Principles)",
	"La Fête des Récompenses (Celebration of Honours)",
	"La Fête de la Révolution (Celebration of the Revolution)",
}
