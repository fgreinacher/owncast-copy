var request = require('supertest');

const Random = require('crypto-random');

const sendConfigChangeRequest = require('./lib/config').sendConfigChangeRequest;
request = request('http://127.0.0.1:8080');

const serverName = randomString();
const streamTitle = randomString();
const serverSummary = randomString();
const offlineMessage = randomString();
const pageContent = `<p>${randomString()}</p>`;
const tags = [randomString(), randomString(), randomString()];
const streamKeys = [
	{ key: randomString(), comment: 'test key 1' },
	{ key: randomString(), comment: 'test key 2' },
	{ key: randomString(), comment: 'test key 3' },
];

const latencyLevel = Math.floor(Math.random() * 4);
const appearanceValues = {
	variable1: randomString(),
	variable2: randomString(),
	variable3: randomString(),
};

const streamOutputVariants = {
	videoBitrate: randomNumber() * 100,
	framerate: 42,
	cpuUsageLevel: 2,
	scaledHeight: randomNumber() * 100,
	scaledWidth: randomNumber() * 100,
};
const socialHandles = [
	{
		url: 'http://facebook.org/' + randomString(),
		platform: randomString(),
	},
];

const s3Config = {
	enabled: true,
	endpoint: 'http://' + randomString() + ".tld",
	accessKey: randomString(),
	secret: randomString(),
	bucket: randomString(),
	region: randomString(),
	forcePathStyle: true,
};

const forbiddenUsernames = [randomString(), randomString(), randomString()];

const ypConfig = {
	enabled: true,
	instanceUrl: 'http://' + randomString() + ".tld"
};

const federationConfig = {
	enabled: true,
	isPrivate: true,
	username: randomString(),
	goLiveMessage: randomString(),
	showEngagement: false,
	blockedDomains: [randomString() + ".tld", randomString() + ".tld"],
};

const defaultAdminPassword = 'abc123';
const defaultStreamKey = undefined;

test('verify default streamKey', async (done) => {
	request
		.get('/api/admin/serverconfig')
		.auth('admin', defaultAdminPassword)
		.expect(200)
		.then((res) => {
			expect(res.body.streamKey).toBe(defaultStreamKey);
			done();
		});
});

test('verify default directory configurations', async (done) => {
	request
		.get('/api/admin/serverconfig')
		.auth('admin', defaultAdminPassword)
		.expect(200)
		.then((res) => {
			expect(res.body.yp.enabled).toBe(!ypConfig.enabled);
			done();
		});
});

test('verify default federation configurations', async (done) => {
	request
		.get('/api/admin/serverconfig')
		.auth('admin', defaultAdminPassword)
		.expect(200)
		.then((res) => {
			expect(res.body.federation.enabled).toBe(!federationConfig.enabled);
			expect(res.body.federation.isPrivate).toBe(!federationConfig.isPrivate);
			expect(res.body.federation.showEngagement).toBe(!federationConfig.showEngagement);
			expect(res.body.federation.goLiveMessage).toBe("I've gone live!");
			expect(res.body.federation.blockedDomains).toStrictEqual([]);
			done();
		});
});

test('set server name', async (done) => {
	const res = await sendConfigChangeRequest('name', serverName);
	done();
});

test('set stream title', async (done) => {
	const res = await sendConfigChangeRequest('streamtitle', streamTitle);
	done();
});

test('set server summary', async (done) => {
	const res = await sendConfigChangeRequest('serversummary', serverSummary);
	done();
});

test('set extra page content', async (done) => {
	const res = await sendConfigChangeRequest('pagecontent', pageContent);
	done();
});

test('set tags', async (done) => {
	const res = await sendConfigChangeRequest('tags', tags);
	done();
});

test('set stream keys', async (done) => {
	const res = await sendConfigChangeRequest('streamkeys', streamKeys);
	done();
});

test('set latency level', async (done) => {
	const res = await sendConfigChangeRequest(
		'video/streamlatencylevel',
		latencyLevel
	);
	done();
});

test('set video stream output variants', async (done) => {
	const res = await sendConfigChangeRequest('video/streamoutputvariants', [
		streamOutputVariants,
	]);
	done();
});

test('set social handles', async (done) => {
	const res = await sendConfigChangeRequest('socialhandles', socialHandles);
	done();
});

test('set s3 configuration', async (done) => {
	const res = await sendConfigChangeRequest('s3', s3Config);
	done();
});

test('set forbidden usernames', async (done) => {
	const res = await sendConfigChangeRequest(
		'chat/forbiddenusernames',
		forbiddenUsernames
	);
	done();
});

test('set server url', async (done) => {
	const res = await sendConfigChangeRequest('serverurl', ypConfig.instanceUrl);
	done();
});

test('set federation username', async (done) => {
	const res = await sendConfigChangeRequest('federation/username', federationConfig.username);
	done();
});

test('set federation goLiveMessage', async (done) => {
	const res = await sendConfigChangeRequest('federation/livemessage', federationConfig.goLiveMessage);
	done();
});

test('set hide viewer count', async (done) => {
	const res = await sendConfigChangeRequest('hideviewercount', true);
	done();
});

test('toggle private federation mode', async (done) => {
	const res = await sendConfigChangeRequest('federation/private', federationConfig.isPrivate);
	done();
});

test('toggle federation engagement', async (done) => {
	const res = await sendConfigChangeRequest('federation/showengagement', federationConfig.showEngagement);
	done();
});

test('set federation blocked domains', async (done) => {
	const res = await sendConfigChangeRequest('federation/blockdomains', federationConfig.blockedDomains);
	done();
});


test('set offline message', async (done) => {
	const res = await sendConfigChangeRequest('offlinemessage', offlineMessage);
	done();
});

test('set custom style values', async (done) => {
	const res = await sendConfigChangeRequest('appearance', appearanceValues);
	done();
});

test('enable directory', async (done) => {
	const res = await sendConfigChangeRequest('directoryenabled', true);
	done();
});

test('enable federation', async (done) => {
	const res = await sendConfigChangeRequest('federation/enable', federationConfig.enabled);
	done();
});

test('verify updated config values', async (done) => {
	const res = await request.get('/api/config');
	expect(res.body.name).toBe(serverName);
	expect(res.body.streamTitle).toBe(streamTitle);
	expect(res.body.summary).toBe(`${serverSummary}`);
	expect(res.body.extraPageContent).toBe(pageContent);
	expect(res.body.offlineMessage).toBe(offlineMessage);
	expect(res.body.logo).toBe('/logo');
	expect(res.body.socialHandles).toStrictEqual(socialHandles);
	done();
});

// Test that the raw video details being broadcasted are coming through
test('admin stream details are correct', (done) => {
	request
		.get('/api/admin/status')
		.auth('admin', defaultAdminPassword)
		.expect(200)
		.then((res) => {
			expect(res.body.broadcaster.streamDetails.width).toBe(320);
			expect(res.body.broadcaster.streamDetails.height).toBe(180);
			expect(res.body.broadcaster.streamDetails.framerate).toBe(24);
			expect(res.body.broadcaster.streamDetails.videoBitrate).toBe(1269);
			expect(res.body.broadcaster.streamDetails.videoCodec).toBe('H.264');
			expect(res.body.broadcaster.streamDetails.audioCodec).toBe('AAC');
			expect(res.body.online).toBe(true);
			done();
		});
});

test('admin configuration is correct', (done) => {
	request
		.get('/api/admin/serverconfig')
		.auth('admin', defaultAdminPassword)
		.expect(200)
		.then((res) => {
			expect(res.body.instanceDetails.name).toBe(serverName);
			expect(res.body.instanceDetails.summary).toBe(serverSummary);
			expect(res.body.instanceDetails.offlineMessage).toBe(offlineMessage);
			expect(res.body.instanceDetails.tags).toStrictEqual(tags);
			expect(res.body.instanceDetails.socialHandles).toStrictEqual(
				socialHandles
			);
			expect(res.body.forbiddenUsernames).toStrictEqual(forbiddenUsernames);
			expect(res.body.streamKeys).toStrictEqual(streamKeys);

			expect(res.body.videoSettings.latencyLevel).toBe(latencyLevel);
			expect(res.body.videoSettings.videoQualityVariants[0].framerate).toBe(
				streamOutputVariants.framerate
			);
			expect(res.body.videoSettings.videoQualityVariants[0].cpuUsageLevel).toBe(
				streamOutputVariants.cpuUsageLevel
			);

			expect(res.body.yp.enabled).toBe(true);
			expect(res.body.yp.instanceUrl).toBe(ypConfig.instanceUrl);

			expect(res.body.adminPassword).toBe(defaultAdminPassword);

			expect(res.body.s3.enabled).toBe(s3Config.enabled);
			expect(res.body.s3.endpoint).toBe(s3Config.endpoint);
			expect(res.body.s3.accessKey).toBe(s3Config.accessKey);
			expect(res.body.s3.secret).toBe(s3Config.secret);
			expect(res.body.s3.bucket).toBe(s3Config.bucket);
			expect(res.body.s3.region).toBe(s3Config.region);
			expect(res.body.s3.forcePathStyle).toBe(true);
			expect(res.body.hideViewerCount).toBe(true);


			expect(res.body.federation.enabled).toBe(federationConfig.enabled);
			expect(res.body.federation.isPrivate).toBe(federationConfig.isPrivate);
			expect(res.body.federation.username).toBe(federationConfig.username);
			expect(res.body.federation.goLiveMessage).toBe(federationConfig.goLiveMessage);
			expect(res.body.federation.showEngagement).toBe(federationConfig.showEngagement);
			expect(res.body.federation.blockedDomains).toStrictEqual(federationConfig.blockedDomains);
			done();
		});
});

test('frontend configuration is correct', (done) => {
	request
		.get('/api/config')
		.expect(200)
		.then((res) => {
			expect(res.body.name).toBe(serverName);
			expect(res.body.logo).toBe('/logo');
			expect(res.body.socialHandles).toStrictEqual(socialHandles);
			done();
		});
});

test('frontend status is correct', (done) => {
	request
		.get('/api/status')
		.expect(200)
		.then((res) => {
			expect(res.body.viewerCount).toBe(undefined);
			done();
		});
});


function randomString(length = 20) {
	return Random.value().toString(16).substr(2, length);
}

function randomNumber() {
	return Random.range(0, 5);
}
