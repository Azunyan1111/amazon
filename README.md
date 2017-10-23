# amazon

## HOW TO
your amazon mws api keys
```model.AmazonAPI.ApiInit()
config := gmws.MwsConfig{
		SellerId:  os.Getenv("SellerId"),
		AccessKey: os.Getenv("AccessKey"),
		SecretKey: os.Getenv("SecretKey"),
		Region:    "JP",
	}
```

Database
your database.
```model.Database.DataBaseInit()
dataSource := os.Getenv("DATABASE_URL")
```
create
```
-- Create syntax for TABLE 'Items'
CREATE TABLE `Items` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ASIN` char(10) NOT NULL DEFAULT '',
  `title` text,
  `image` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ASIN` (`ASIN`)
) ENGINE=InnoDB AUTO_INCREMENT=1863 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'CategoryURL'
CREATE TABLE `CategoryURL` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `URL` char(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `urlIndex` (`URL`)
) ENGINE=InnoDB AUTO_INCREMENT=21292 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'Price'
CREATE TABLE `Price` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ASIN` char(10) NOT NULL DEFAULT '',
  `Amount` char(13) NOT NULL DEFAULT '',
  `Channel` char(255) DEFAULT '',
  `Conditions` char(255) DEFAULT NULL,
  `ShippingTime` char(255) DEFAULT NULL,
  `InsertTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1232 DEFAULT CHARSET=utf8;
```

# used

[gomws](https://github.com/svvu/gomws)
[github.com/juju/errors](https://github.com/juju/errors)
